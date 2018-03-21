module.exports = {
  entry: [
    "./src/index.js",
  ],
  output: {
    path: __dirname,
    filename: "./src/bundle.js"
  },
  node: {
    fs: 'empty'
  }
};
