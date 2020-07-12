const path = require('path');

module.exports = {
  mode: 'production',
  target: 'web',
  entry: {
    ssr: path.resolve(__dirname, './main.tsx'),
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    library: 'SSR',
    libraryTarget: 'var',
    filename: '[name].js'
  },
  resolve: {
    extensions: [".js", ".jsx", ".json", ".ts", ".tsx"]
  },
  module: {
    rules: [
      {
        test: /\.ts(x?)$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "ts-loader"
          }
        ]
      }
    ]
  }
};