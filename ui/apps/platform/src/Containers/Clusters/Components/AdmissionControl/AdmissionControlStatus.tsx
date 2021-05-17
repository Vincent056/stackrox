import React, { ReactElement } from 'react';

import { Tooltip, DetailedTooltipOverlay } from '@stackrox/ui-components';

import { getDistanceStrictAsPhrase } from 'utils/dateUtils';
import HealthStatus from '../HealthStatus';
import {
    delayedAdmissionControlStatusStyle,
    healthStatusStyles,
    isDelayedSensorHealthStatus,
} from '../../cluster.helpers';
import AdmissionControlStatusTotals from './AdmissionControlStatusTotals';
import AdmissionControlUnavailableStatus from './AdmissionControlUnavailableStatus';
import { ClusterHealthStatus } from '../../clusterTypes';
import HealthLabelWithDelayed from '../HealthLabelWithDelayed';
import HealthStatusNotApplicable from '../HealthStatusNotApplicable';

/*
 * Admission Control Status in Clusters list if `isList={true}` or Cluster side panel if `isList={false}`
 *
 * Caller is responsible for optional chaining in case healthStatus is null.
 */

type AdmissionControlStatusProps = {
    healthStatus: ClusterHealthStatus;
    isList?: boolean;
};

function AdmissionControlStatus({
    healthStatus,
    isList = false,
}: AdmissionControlStatusProps): ReactElement {
    if (!healthStatus?.admissionControlHealthStatus) {
        return <HealthStatusNotApplicable testId="admissionControlStatus" />;
    }

    const {
        admissionControlHealthStatus,
        admissionControlHealthInfo,
        healthInfoComplete,
        sensorHealthStatus,
        lastContact,
    } = healthStatus;
    const isDelayed = !!(lastContact && isDelayedSensorHealthStatus(sensorHealthStatus));
    const { Icon, bgColor, fgColor } = isDelayed
        ? delayedAdmissionControlStatusStyle
        : healthStatusStyles[admissionControlHealthStatus];
    const icon = <Icon className="h-4 w-4" />;
    const currentDatetime = new Date();

    const healthLabelElement = (
        <HealthLabelWithDelayed
            isDelayed={isDelayed}
            delayedText={getDistanceStrictAsPhrase(lastContact, currentDatetime)}
            clusterHealthItem="admissionControl"
            clusterHealthItemStatus={admissionControlHealthStatus}
            isList={isList}
        />
    );

    const healthStatusElement = (
        <HealthStatus icon={icon} iconColor={fgColor}>
            {healthLabelElement}
        </HealthStatus>
    );

    if (admissionControlHealthInfo) {
        const admissionControlTotalsElement = (
            <AdmissionControlStatusTotals
                bgColor={bgColor}
                fgColor={fgColor}
                admissionControlHealthInfo={admissionControlHealthInfo}
            />
        );
        const infoElement = healthInfoComplete ? (
            admissionControlTotalsElement
        ) : (
            <div>
                {admissionControlTotalsElement}
                <div data-testid="admissionControlInfoComplete">
                    <strong>Upgrade Sensor</strong> to get complete Admission Control health
                    information
                </div>
            </div>
        );

        return isList ? (
            <Tooltip
                content={
                    <DetailedTooltipOverlay
                        title="Admission Control Health Information"
                        body={infoElement}
                    />
                }
            >
                <div>{healthStatusElement}</div>
            </Tooltip>
        ) : (
            <HealthStatus icon={icon} iconColor={fgColor}>
                <div>
                    {healthLabelElement}
                    {infoElement}
                </div>
            </HealthStatus>
        );
    }

    if (admissionControlHealthStatus === 'UNAVAILABLE') {
        return (
            <AdmissionControlUnavailableStatus
                isList={isList}
                icon={icon}
                fgColor={fgColor}
                healthStatusElement={healthStatusElement}
                healthLabelElement={healthLabelElement}
            />
        );
    }

    // UNINITIALIZED
    return healthStatusElement;
}

export default AdmissionControlStatus;
