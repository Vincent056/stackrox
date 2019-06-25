package services

import static com.jayway.restassured.RestAssured.given

import com.jayway.restassured.config.RestAssuredConfig
import util.Keys

import javax.net.ssl.SSLContext
import java.security.SecureRandom

import com.jayway.restassured.config.SSLConfig
import org.apache.http.conn.ssl.SSLSocketFactory
import util.Env

import io.stackrox.proto.api.v1.AuthProviderServiceGrpc
import io.stackrox.proto.api.v1.AuthproviderService
import io.stackrox.proto.api.v1.Common
import io.stackrox.proto.storage.AuthProviderOuterClass

class AuthProviderService extends BaseService {
    static getAuthProviderService() {
        return AuthProviderServiceGrpc.newBlockingStub(getChannel())
    }

    static getAuthProviders() {
        return getAuthProviderService().getAuthProviders(
                AuthproviderService.GetAuthProvidersRequest.newBuilder().build()
        )
    }

    static getAuthProvider(String id) {
        try {
            return getAuthProviderService().getAuthProvider(
                    AuthproviderService.GetAuthProviderRequest.newBuilder().setId(id).build()
            )
        } catch (Exception e) {
            println "Failed getting auth provider: ${e.toString()}"
        }
    }

    static createAuthProvider(String name, String type, Map<String, String> config) {
        try {
            return getAuthProviderService().postAuthProvider(
                    AuthproviderService.PostAuthProviderRequest.newBuilder().setProvider(
                            AuthProviderOuterClass.AuthProvider.newBuilder()
                                    .setName(name)
                                    .setType(type)
                                    .putAllConfig(config)
                                    .setEnabled(true)
                    ).build()
            ).id
        } catch (Exception e) {
            println "Failed to create auth provider: ${e.toString()}"
        }
    }

    static deleteAuthProvider(String id) {
        getAuthProviderService().deleteAuthProvider(Common.ResourceByID.newBuilder().setId(id).build())
    }

    static getAuthProviderLoginToken(String id) {
        String loginUrl = getAuthProvider(id).loginUrl

        def sslContext = SSLContext.getInstance("TLS")
        sslContext.init(Keys.keyManagerFactory().keyManagers, Keys.trustManagerFactory().trustManagers,
                new SecureRandom())

        def socketFactory = new SSLSocketFactory(sslContext, SSLSocketFactory.ALLOW_ALL_HOSTNAME_VERIFIER)

        def location = loginUrl
        // There are two redirects: first from the generic URL to the auth provider's URL, and then from the auth
        // provider's URL to the token response URL.
        for (int i = 0; i < 2; i++) {
            def response =
                    given().header("Content-Type", "application/json")
                            .config(RestAssuredConfig.newConfig().sslConfig(
                            SSLConfig.sslConfig().with().sslSocketFactory(socketFactory)
                                    .and().allowAllHostnames()))
                            .when()
                            .redirects().follow(false)
                            .get("https://${Env.mustGetHostname()}:${Env.mustGetPort()}${location}")
            location = response.getHeader("Location")
        }
        def fullURL = new URL("https://${Env.mustGetHostname()}:${Env.mustGetPort()}${location}")
        def token = ""
        fullURL.ref.split("&").each {
            def values = it.split("=")
            if (values[0] == "token") {
                token = values[1]
            }
        }
        assert token != "" : "Could not determine token for cert"
        return token
    }
}
