/**
 * Created by alexlivenson1 on 3/31/16.
 */
import React from 'react';

export default React.createClass({
    render: function () {
        return (
            <div className="greeting">
                Hello, {this.props.name}!
            </div>
        );
    }
})
