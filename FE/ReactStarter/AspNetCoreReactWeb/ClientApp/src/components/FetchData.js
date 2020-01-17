import React, { Component } from 'react';
import { Table } from 'antd';

export class FetchData extends Component {
  displayName = FetchData.name

  constructor(props) {
    super(props);
    this.state = {
      data: [], 
      columns: [{
        title: 'temperatureF',
        dataIndex: 'temperatureF',
        key: 'temperatureF',
      },{
        title: 'temperatureC',
        dataIndex: 'temperatureC',
        key: 'temperatureC',
      }, {
        title: 'summary',
        dataIndex: 'summary',
        key: 'summary',
      }, {
        title: 'dateFormatted',
        dataIndex: 'dateFormatted',
        key: 'dateFormatted',
      }],
      loading: true
    };

    fetch('api/SampleData/WeatherForecasts')
      .then(response => response.json())
      .then(data => {
        console.info(data);

        this.setState({ data: data, loading: false });
      });
  }

  render() {
    return (
      <div>
        <h1>Weather forecast</h1>
        <p>This component demonstrates fetching data from the server.</p>
        <Table
        columns={this.state.columns}
        rowKey={record => record.id}
        dataSource={this.state.data}
        loading={this.state.loading}
        onChange={this.handleTableChange}
      />
      </div>
    );
  }
}
