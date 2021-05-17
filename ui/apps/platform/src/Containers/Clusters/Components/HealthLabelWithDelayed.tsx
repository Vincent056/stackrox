import React, { ReactElement } from 'react';

import { healthStatusLabels } from 'messages/common';
import {
    delayedCollectorStatusStyle,
    delayedAdmissionControlStatusStyle,
    healthStatusStyles,
} from '../cluster.helpers';
import { ClusterHealthItemStatus, ClusterHealthItem } from '../clusterTypes';

type HealthLabelWithDelayedProps = {
    delayedText: string;
    clusterHealthItem: ClusterHealthItem;
    clusterHealthItemStatus: ClusterHealthItemStatus;
    isDelayed: boolean;
    isList: boolean;
};

// In rare case that the block does not fit in a narrow column,
// the space and "whitespace-nowrap" cause time phrase to wrap as a unit.
// Order arguments according to date-fns@2 convention:
// If lastContact <= currentDateTime: X units ago
function HealthLabelWithDelayed({
    isDelayed,
    delayedText,
    clusterHealthItem,
    clusterHealthItemStatus,
    isList,
}: HealthLabelWithDelayedProps): ReactElement {
    let { bgColor, fgColor } = healthStatusStyles[clusterHealthItemStatus];
    if (isDelayed) {
        if (clusterHealthItem === 'collector') {
            bgColor = delayedCollectorStatusStyle.bgColor;
            fgColor = delayedCollectorStatusStyle.fgColor;
        }
        if (clusterHealthItem === 'admissionControl') {
            bgColor = delayedAdmissionControlStatusStyle.bgColor;
            fgColor = delayedAdmissionControlStatusStyle.fgColor;
        }
    }
    const testId = `${clusterHealthItem}Status`;
    const healthLabelText = isList
        ? clusterHealthItem
        : healthStatusLabels[clusterHealthItemStatus];
    const healthLabelElement = (
        <span className={`${bgColor} ${fgColor} capitalize`}>{healthLabelText}</span>
    );
    if (isDelayed) {
        return (
            <div data-testid={testId}>
                {healthLabelElement}
                <span className="whitespace-nowrap">{` ${delayedText}`}</span>
            </div>
        );
    }
    return <div data-testid={testId}>{healthLabelElement}</div>;
}

export default HealthLabelWithDelayed;
