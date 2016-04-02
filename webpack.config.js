// Reference: https://robots.thoughtbot.com/setting-up-webpack-for-react-and-hot-module-replacement
var webpack = require('webpack');

module.exports = {
    context: __dirname + "/public",
    /*
     Webpack needs to copy this file over to the `dist` folder to us to use => need to modify the `entry` property
     in the webpack.
     */
    entry: {
        javascript: "./js/index.js",
        html: "./views/index.html"
    },
    output: {
        filename: "bundle.js",
        path: __dirname + "/dist"
    },
    plugins: [
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery"
        })
    ],
    /*
     Webpack accepts an array of loader objects which specify loaders to apply to files that match the test regex
     and exclude files that match `exclude` regex. In this case, applying `babel-loader` to all files with .js
     ext not in node_modules
     */
    module: {
        loaders: [{
            test: /\.js$/,
            exclude: /node_modules/,
            loader: "babel-loader",
            query: {
                presets: ['es2015', 'react']
            }
        }, {
            test: /\.html$/,
            loader: "file?name=[name].[ext]"
        }, {
            test: /\.css$/,
            loader: "style!css!" // style!css! => take css and pipe to css loader and then to style loader
        }]
    }
};