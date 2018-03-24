var webpack = require('webpack');

module.exports = {
  mode: 'development',
  context: __dirname,
  entry: [
    'webpack-hot-middleware/client?path=/__webpack_hmr&timeout=20000',
    "./src/index.js"
  ],
  devtool: '#source-map',
  plugins: [
    new webpack.optimize.OccurrenceOrderPlugin(),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoEmitOnErrorsPlugin()
  ],
  output: {
    path: __dirname,
    publicPath: '/',
    filename: "./src/bundle.js"
  },
  node: {
    fs: 'empty'
  }
};