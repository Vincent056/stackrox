type Directionality = 'egress' | 'ingress';

type Protocols = 'L4_PROTOCOL_TCP' | 'L4_PROTOCOL_UDP';

type Ports = string; // number format

export type FilterValue = Directionality | Protocols | Ports;

export type AdvancedFlowsFilterType = {
    directionality: Directionality[];
    protocols: Protocols[];
    ports: Ports[];
};
