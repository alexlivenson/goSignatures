/**
 * Created by alexlivenson1 on 3/30/16.
 */
/*
 This is used by webpack as an entry point to load all the files
 */
'use strict';

import React from 'react';
import ReactDom from 'react-dom';
import '../styles/style.css';
import 'bootstrap';

var FancyButton = React.createClass({
    render: function () {
        return <button onClick={this.props.onClick}>
            <i className={"fa " + this.props.icon}/>
            <span>{this.props.text}</span>
        </button>
    }
});

class HelloWorld extends React.Component {
    constructor(props) {
        super(props);
        this.state = {counter: 0};
    }

    increment() {
        this.setState({counter: this.state.counter + 1});
    }

    render() {
        return <div>
            <div className="btn-danger">{this.state.counter}</div>
            <FancyButton text="Increment Baby!" icon="fa-arrow-circle-o-up" onClick={this.increment.bind(this)}/>
        </div>;
    }
}


ReactDom.render(
    <HelloWorld />, document.body
);