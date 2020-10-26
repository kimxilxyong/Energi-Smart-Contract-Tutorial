/* eslint-disable ParseError */
// svelte.config.js - NOTE: you cannot use the new "import x from y" and "export const" syntax in here.
const sveltePreprocess = require('svelte-preprocess');

const preprocessOptions = {
    sourceMap: false, // "you would always want sourcemaps for the IDE" – dummdidumm
};

module.exports = {
    preprocess: sveltePreprocess(preprocessOptions),
    // ...other svelte options
};
