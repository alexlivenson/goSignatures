/**
 * Created by alexlivenson1 on 3/31/16.
 */
import React from 'react';

class Greeting extends React.Component {
    render() {
        return (
            <div className="greeting">
                Hello, {this.props.name}!
            </div>
        );
    }
}

export default Greeting