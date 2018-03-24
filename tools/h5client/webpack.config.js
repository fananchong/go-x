var webpack = require('webpack');
var hotMiddlewareScript = 'webpack-hot-middleware/client?path=/__webpack_hmr&timeout=20000&reload=false';

module.exports = {
  mode: 'development',
  context: __dirname,
  entry: {
    index: ["./src/index.js", hotMiddlewareScript]
  },
  devtool: '#source-map',
  plugins: [
    new webpack.optimize.OccurrenceOrderPlugin(),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoEmitOnErrorsPlugin()
  ],
  output: {
    path: __dirname,
    publicPath: 'http://localhost:8000/',
    filename: './src/bundle.js'
  },
  devServer: {
    "wwwPath": './src',
    "open": true,
    "browser": ["chrome", '--allow-file-access-from-files', '--disable-web-security', '--user-data-dir=./userdata']
  },
  node: {
    fs: 'empty'
  }
};