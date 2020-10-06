/* eslint-disable ParseError */
// svelte.config.js - NOTE: you cannot use the new "import x from y" and "export const" syntax in here.
const sveltePreprocess = require('svelte-preprocess');

const preprocessOptions = {
    sourceMap: true, // "you would always want sourcemaps for the IDE" – dummdidumm
    defaults: {
        script: 'javascript', // <-- now you can just write <script>let typingsAllowed: string;</script>
    },
};

module.exports = {
    preprocess: sveltePreprocess(preprocessOptions),
    // ...other svelte options
};
