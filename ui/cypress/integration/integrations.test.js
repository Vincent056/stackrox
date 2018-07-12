import { selectors } from './constants/IntegrationsPage';
import * as api from './constants/apiEndpoints';

describe('Integrations page', () => {
    beforeEach(() => {
        cy.visit('/');
        cy.get(selectors.configure).click();
        cy.get(selectors.navLink).click();
    });

    it('Plugin tiles should all be the same height', () => {
        let value = null;
        cy.get(selectors.plugins).each($el => {
            if (value) expect($el[0].clientHeight).to.equal(value);
            else value = $el[0].clientHeight;
        });
    });

    it('should have selected item in nav bar', () => {
        cy.get(selectors.configure).should('have.class', 'bg-primary-600');
    });

    it('should allow integration with Slack', () => {
        cy.get('div.ReactModalPortal').should('not.exist');

        cy.get(selectors.slackTile).click();
        cy.get('div.ReactModalPortal');
    });

    it('should add an integration with DockerHub', () => {
        cy.get(selectors.dockerRegistryTile).click();
        cy.get(selectors.buttons.add).click();

        const name = `Docker Registry ${Math.random()
            .toString(36)
            .substring(7)}`;
        cy.get(selectors.dockerRegistryForm.nameInput).type(name);

        cy.get(`${selectors.dockerRegistryForm.typesSelect} .Select-arrow`).click();
        cy
            .get(
                `${
                    selectors.dockerRegistryForm.typesSelect
                } div[role="option"]:contains("Registry")`
            )
            .click();

        // test that validation error happens when form is incomplete
        cy.get(selectors.buttons.test).click();
        cy.get(selectors.integrationError);

        cy.get(selectors.dockerRegistryForm.endpointInput).type('registry-1.docker.io');

        cy.get(selectors.buttons.create).click();

        // delete the integration after to clean up
        cy.get(`table tr:contains("${name}") td input[type="checkbox"]`).check();
        cy.get(selectors.buttons.delete).click();
        cy.get(selectors.buttons.confirm).click();
        cy.get(`table tr:contains("${name}")`).should('not.exist');
    });
});

describe('Cluster Creation Flow', () => {
    beforeEach(() => {
        cy.server();
        cy.fixture('clusters/single.json').as('singleCluster');
        cy.route('GET', api.clusters.list, '@singleCluster').as('clusters');
        cy.route('POST', api.clusters.zip, {}).as('download');
        cy.route('POST', api.clusters.list).as('addCluster');
        cy.visit('/');
        cy.get(selectors.configure).click();
        cy.get(selectors.navLink).click();
        cy.wait('@clusters');
    });

    it('Should show a confirmation dialog when trying to delete clusters', () => {
        cy.get(selectors.dockerSwarmTile).click();
        cy.get(selectors.dialog).should('not.exist');
        cy.get(selectors.checkboxes).check();
        cy.get(selectors.buttons.delete).click();
        cy.get(selectors.dialog);
    });

    it('Should show the remote cluster when clicking the Docker Swarm tile', () => {
        cy.get(selectors.dockerSwarmTile).click();

        cy.get(selectors.clusters.swarmCluster1);
    });

    it('Should show a disabled form when viewing a specific cluster', () => {
        cy.get(selectors.dockerSwarmTile).click();

        cy.get(selectors.clusters.swarmCluster1).click();

        cy
            .get(selectors.readOnlyView)
            .eq(0)
            .should('have.text', 'Name:Swarm Cluster 1');

        cy
            .get(selectors.readOnlyView)
            .eq(1)
            .should('have.text', 'Cluster Type:Swarm');

        cy
            .get(selectors.readOnlyView)
            .eq(2)
            .should('have.text', 'Prevent Image:stackrox/prevent:latest');

        cy
            .get(selectors.readOnlyView)
            .eq(3)
            .should('have.text', 'Central API Endpoint:central.stackrox:443');
    });

    it.only('Should be able to fill out the Swarm form, download config files and see cluster checked-in', () => {
        cy.get(selectors.dockerSwarmTile).click();

        cy.get(selectors.buttons.add).click();

        const clusterName = 'Swarm Cluster TestInstance';
        cy.get(selectors.clusterForm.nameInput).type(clusterName);
        cy.get(selectors.clusterForm.imageInput).type('stackrox/prevent:latest');
        cy.get(selectors.clusterForm.endpointInput).type('central.prevent_net:443');

        cy.get(selectors.buttons.next).click();
        cy
            .wait('@addCluster')
            .its('responseBody')
            .then(response => {
                const clusterId = response.cluster.id;

                cy.get(selectors.buttons.download).click();
                cy.wait('@download');

                cy.get('div:contains("Waiting for the cluster to check-in successfully...")');

                // make cluster to "check-in" by adding "lastContact"
                cy
                    .route('GET', `${api.clusters.list}/${clusterId}`, {
                        cluster: {
                            id: clusterId,
                            lastContact: '2018-06-25T19:12:44.955289Z'
                        }
                    })
                    .as('getCluster');
                cy.wait('@getCluster');
                cy.get(
                    'div:contains("Success! The cluster has been recognized properly by Prevent. You may now save the configuration.")'
                );

                // clean up after the test by deleting the cluster
                cy.get(`table tr:contains("${clusterName}") td input[type="checkbox"]`).check();
                cy.get(selectors.buttons.add).should('be.disabled');
                cy.get(selectors.buttons.delete).click();
                cy.get(selectors.buttons.confirm).click();
                cy.get(`table tr:contains("${clusterName}")`).should('not.exist');
            });
    });
});
