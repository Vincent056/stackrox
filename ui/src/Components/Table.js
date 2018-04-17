import React, { Component } from 'react';
import PropTypes from 'prop-types';
import isEqual from 'lodash/isEqual';

import TableCell from 'Components/TableCell';
import find from 'lodash/find';
import flattenObject from 'utils/flattenObject';

class Table extends Component {
    static propTypes = {
        columns: PropTypes.arrayOf(
            PropTypes.shape({
                key: PropTypes.string,
                label: PropTypes.string,
                keyValueFunc: PropTypes.func,
                align: PropTypes.string,
                classFunc: PropTypes.func,
                className: PropTypes.string,
                default: PropTypes.any,
                sortMethod: PropTypes.func
            })
        ).isRequired,
        rows: PropTypes.arrayOf(
            PropTypes.shape({
                id: PropTypes.string
            })
        ).isRequired,
        onRowClick: PropTypes.func,
        onRowChecked: PropTypes.func,
        checkboxes: PropTypes.bool,
        actions: PropTypes.arrayOf(
            PropTypes.shape({
                text: PropTypes.string,
                renderIcon: PropTypes.func,
                className: PropTypes.string,
                onClick: PropTypes.func,
                disabled: PropTypes.bool
            })
        )
    };

    static defaultProps = {
        onRowClick: null,
        onRowChecked: null,
        checkboxes: false,
        actions: []
    };

    constructor(props) {
        super(props);

        this.state = {
            sortBy: null,
            sortDir: {},
            checked: new Set(),
            selected: null,
            sortedRows: new Set()
        };
    }

    componentWillMount() {
        this.setState({ sortedRows: this.props.rows });
    }

    componentWillReceiveProps = props => {
        if (props.rows !== this.props.rows) {
            this.setState({ sortedRows: props.rows });
        }
    };

    getSelectedRows = () => Array.from(this.state.checked);

    getDirection = direction => {
        if (!direction) return '';
        return direction === 'DESC' ? ' ↓' : ' ↑';
    };

    getValue = (obj, key) => {
        let val = Object.assign({}, flattenObject(obj))[key];
        if (val) {
            val = typeof val === 'string' ? val : val[0];
        }
        return val;
    };

    clearSelectedRows = () => {
        const { checked } = this.state;
        checked.clear();
        this.setState({ checked });
    };

    rowCheckedHandler = row => event => {
        event.stopPropagation();
        const { checked } = this.state;
        if (!checked.has(row)) checked.add(row);
        else checked.delete(row);
        this.setState({ checked });
        if (this.props.onRowChecked) this.props.onRowChecked(Array.from(this.state.checked));
    };

    rowClickHandler = row => () => {
        if (this.props.onRowClick) {
            this.props.onRowClick(row);
            this.setState({
                selected: row
            });
        }
    };

    actionClickHandler = (action, row) => event => {
        event.stopPropagation();
        action.onClick(row);
    };

    sortRows = key => () => {
        const sortBy = key;
        let sortDir = this.state.sortDir[sortBy];

        if (sortBy === this.state.sortBy) {
            sortDir = this.state.sortDir[sortBy] === 'ASC' ? 'DESC' : 'ASC';
        } else {
            sortDir = 'DESC';
        }

        const column = find(this.props.columns, o => o.key === sortBy);
        const sortFn = (a, b) => {
            let sortVal = 0;
            if (column && column.sortMethod) {
                sortVal = column.sortMethod(a, b);
            } else {
                const aValue = this.getValue(a, sortBy);
                const bValue = this.getValue(b, sortBy);
                if (aValue === bValue) sortVal = 0;
                if (aValue === undefined) sortVal = -1;
                if (bValue === undefined) sortVal = 1;
                if (aValue !== undefined && bValue !== undefined)
                    sortVal = aValue.localeCompare(bValue);
            }
            if (sortDir === 'DESC') {
                sortVal *= -1;
            }
            return sortVal;
        };
        const sortedRows = [...this.state.sortedRows].sort(sortFn);
        this.setState({ sortBy, sortDir: { [key]: sortDir }, sortedRows });
    };

    renderActionButtons = row =>
        this.props.actions.map((button, i) => (
            <button
                key={i}
                className={button.className}
                onClick={this.actionClickHandler(button, row)}
                disabled={button.disabled}
            >
                {button.renderIcon && (
                    <span className="flex items-center">{button.renderIcon(row)}</span>
                )}
                {button.text && (
                    <span className={`${button.renderIcon && 'ml-3'}`}>{button.text}</span>
                )}
            </button>
        ));

    renderHeaders() {
        const tableHeaders = this.props.columns.map(column => {
            const className = `p-3 text-primary-500 border-b border-base-300 hover:text-primary-600 cursor-pointer truncate ${
                column.align === 'right' ? 'text-right' : 'text-left'
            } ${column.className}`;
            const key = column.label;
            return (
                <th className={className} key={key} onClick={this.sortRows(column.key)}>
                    {column.label + this.getDirection(this.state.sortDir[column.key])}
                </th>
            );
        });
        if (this.props.checkboxes) {
            tableHeaders.unshift(
                <th
                    className="p-3 text-primary-500 border-b border-base-300 hover:text-primary-600"
                    key="checkboxTableHeader"
                />
            );
        }
        if (this.props.actions && this.props.actions.length) {
            tableHeaders.push(
                <th
                    className="p-3 text-primary-500 border-b border-base-300 hover:text-primary-600"
                    key="actionsTableHeader"
                >
                    Actions
                </th>
            );
        }
        return <tr>{tableHeaders}</tr>;
    }

    renderBody() {
        const { columns } = this.props;
        const rowClickable = !!this.props.onRowClick;
        return [...this.state.sortedRows].map((row, i) => {
            const tableCells = columns.map(column => {
                const key = column.key || column.keys.join('-');
                return <TableCell column={column} row={row} key={`${key}`} />;
            });
            if (this.props.checkboxes) {
                tableCells.unshift(
                    <td className="p-3 text-center" key="checkboxTableCell">
                        <input
                            type="checkbox"
                            className="h-4 w-4 cursor-pointer"
                            onClick={this.rowCheckedHandler(row)}
                            checked={this.state.checked.has(row)}
                        />
                    </td>
                );
            }
            if (this.props.actions && this.props.actions.length) {
                tableCells.push(
                    <td className="flex justify-center p-3 text-center" key="actionsTableCell">
                        {this.renderActionButtons(row)}
                    </td>
                );
            }
            return (
                <tr
                    className={`${rowClickable ? 'cursor-pointer' : ''} ${
                        isEqual(this.state.selected, row) ? 'bg-base-200' : ''
                    } border-b border-base-300 hover:bg-base-100`}
                    key={i}
                    onClick={rowClickable ? this.rowClickHandler(row) : null}
                >
                    {tableCells}
                </tr>
            );
        });
    }

    render() {
        return (
            <table className="w-full border-collapse transition">
                <thead>{this.renderHeaders()}</thead>
                <tbody>{this.renderBody()}</tbody>
            </table>
        );
    }
}

export default Table;
