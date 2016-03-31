// Reference: https://robots.thoughtbot.com/setting-up-webpack-for-react-and-hot-module-replacement
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
        }]
    },
    output: {
        filename: "index.js",
        path: __dirname + "/dist"
    }
};