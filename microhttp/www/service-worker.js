!function(e){var t={};function n(s){if(t[s])return t[s].exports;var r=t[s]={i:s,l:!1,exports:{}};return e[s].call(r.exports,r,r.exports,n),r.l=!0,r.exports}n.m=e,n.c=t,n.d=function(e,t,s){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:s})},n.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var s=Object.create(null);if(n.r(s),Object.defineProperty(s,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)n.d(s,r,function(t){return e[t]}.bind(null,r));return s},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="",n(n.s=2)}([function(e,t,n){"use strict";try{self["workbox:precaching:5.1.4"]&&_()}catch(e){}},function(e,t,n){"use strict";try{self["workbox:core:5.1.4"]&&_()}catch(e){}},function(e,t,n){"use strict";n.r(t);n(0);const s=[],r={get:()=>s,add(e){s.push(...e)}};n(1);const c={googleAnalytics:"googleAnalytics",precache:"precache-v2",prefix:"workbox",runtime:"runtime",suffix:"undefined"!=typeof registration?registration.scope:""},a=e=>[c.prefix,e,c.suffix].filter(e=>e&&e.length>0).join("-"),o=e=>e||a(c.precache),i=(e,...t)=>{let n=e;return t.length>0&&(n+=" :: "+JSON.stringify(t)),n};class h extends Error{constructor(e,t){super(i(e,t)),this.name=e,this.details=t}}const l=new Set;const u=(e,t)=>e.filter(e=>t in e),f=async({request:e,mode:t,plugins:n=[]})=>{const s=u(n,"cacheKeyWillBeUsed");let r=e;for(const e of s)r=await e.cacheKeyWillBeUsed.call(e,{mode:t,request:r}),"string"==typeof r&&(r=new Request(r));return r},d=async({cacheName:e,request:t,event:n,matchOptions:s,plugins:r=[]})=>{const c=await self.caches.open(e),a=await f({plugins:r,request:t,mode:"read"});let o=await c.match(a,s);for(const t of r)if("cachedResponseWillBeUsed"in t){const r=t.cachedResponseWillBeUsed;o=await r.call(t,{cacheName:e,event:n,matchOptions:s,cachedResponse:o,request:a})}return o},p=async({cacheName:e,request:t,response:n,event:s,plugins:r=[],matchOptions:c})=>{const a=await f({plugins:r,request:t,mode:"write"});if(!n)throw new h("cache-put-with-no-response",{url:(o=a.url,new URL(String(o),location.href).href.replace(new RegExp("^"+location.origin),""))});var o;const i=await(async({request:e,response:t,event:n,plugins:s=[]})=>{let r=t,c=!1;for(const t of s)if("cacheWillUpdate"in t){c=!0;const s=t.cacheWillUpdate;if(r=await s.call(t,{request:e,response:r,event:n}),!r)break}return c||(r=r&&200===r.status?r:void 0),r||null})({event:s,plugins:r,response:n,request:a});if(!i)return void 0;const p=await self.caches.open(e),y=u(r,"cacheDidUpdate"),w=y.length>0?await d({cacheName:e,matchOptions:c,request:a}):null;try{await p.put(a,i)}catch(e){throw"QuotaExceededError"===e.name&&await async function(){for(const e of l)await e()}(),e}for(const t of y)await t.cacheDidUpdate.call(t,{cacheName:e,event:s,oldResponse:w,newResponse:i,request:a})},y=async({request:e,fetchOptions:t,event:n,plugins:s=[]})=>{if("string"==typeof e&&(e=new Request(e)),n instanceof FetchEvent&&n.preloadResponse){const e=await n.preloadResponse;if(e)return e}const r=u(s,"fetchDidFail"),c=r.length>0?e.clone():null;try{for(const t of s)if("requestWillFetch"in t){const s=t.requestWillFetch,r=e.clone();e=await s.call(t,{request:r,event:n})}}catch(e){throw new h("plugin-error-request-will-fetch",{thrownError:e})}const a=e.clone();try{let r;r="navigate"===e.mode?await fetch(e):await fetch(e,t);for(const e of s)"fetchDidSucceed"in e&&(r=await e.fetchDidSucceed.call(e,{event:n,request:a,response:r}));return r}catch(e){0;for(const t of r)await t.fetchDidFail.call(t,{error:e,event:n,originalRequest:c.clone(),request:a.clone()});throw e}};let w;async function g(e,t){const n=e.clone(),s={headers:new Headers(n.headers),status:n.status,statusText:n.statusText},r=t?t(s):s,c=function(){if(void 0===w){const e=new Response("");if("body"in e)try{new Response(e.body),w=!0}catch(e){w=!1}w=!1}return w}()?n.body:await n.blob();return new Response(c,r)}function m(e){if(!e)throw new h("add-to-cache-list-unexpected-type",{entry:e});if("string"==typeof e){const t=new URL(e,location.href);return{cacheKey:t.href,url:t.href}}const{revision:t,url:n}=e;if(!n)throw new h("add-to-cache-list-unexpected-type",{entry:e});if(!t){const e=new URL(n,location.href);return{cacheKey:e.href,url:e.href}}const s=new URL(n,location.href),r=new URL(n,location.href);return s.searchParams.set("__WB_REVISION__",t),{cacheKey:s.href,url:r.href}}class R{constructor(e){this._cacheName=o(e),this._urlsToCacheKeys=new Map,this._urlsToCacheModes=new Map,this._cacheKeysToIntegrities=new Map}addToCacheList(e){const t=[];for(const n of e){"string"==typeof n?t.push(n):n&&void 0===n.revision&&t.push(n.url);const{cacheKey:e,url:s}=m(n),r="string"!=typeof n&&n.revision?"reload":"default";if(this._urlsToCacheKeys.has(s)&&this._urlsToCacheKeys.get(s)!==e)throw new h("add-to-cache-list-conflicting-entries",{firstEntry:this._urlsToCacheKeys.get(s),secondEntry:e});if("string"!=typeof n&&n.integrity){if(this._cacheKeysToIntegrities.has(e)&&this._cacheKeysToIntegrities.get(e)!==n.integrity)throw new h("add-to-cache-list-conflicting-integrities",{url:s});this._cacheKeysToIntegrities.set(e,n.integrity)}if(this._urlsToCacheKeys.set(s,e),this._urlsToCacheModes.set(s,r),t.length>0){const e=`Workbox is precaching URLs without revision info: ${t.join(", ")}\nThis is generally NOT safe. Learn more at https://bit.ly/wb-precache`;console.warn(e)}}}async install({event:e,plugins:t}={}){const n=[],s=[],r=await self.caches.open(this._cacheName),c=await r.keys(),a=new Set(c.map(e=>e.url));for(const[e,t]of this._urlsToCacheKeys)a.has(t)?s.push(e):n.push({cacheKey:t,url:e});const o=n.map(({cacheKey:n,url:s})=>{const r=this._cacheKeysToIntegrities.get(n),c=this._urlsToCacheModes.get(s);return this._addURLToCache({cacheKey:n,cacheMode:c,event:e,integrity:r,plugins:t,url:s})});await Promise.all(o);return{updatedURLs:n.map(e=>e.url),notUpdatedURLs:s}}async activate(){const e=await self.caches.open(this._cacheName),t=await e.keys(),n=new Set(this._urlsToCacheKeys.values()),s=[];for(const r of t)n.has(r.url)||(await e.delete(r),s.push(r.url));return{deletedURLs:s}}async _addURLToCache({cacheKey:e,url:t,cacheMode:n,event:s,plugins:r,integrity:c}){const a=new Request(t,{integrity:c,cache:n,credentials:"same-origin"});let o,i=await y({event:s,plugins:r,request:a});for(const e of r||[])"cacheWillUpdate"in e&&(o=e);if(!(o?await o.cacheWillUpdate({event:s,request:a,response:i}):i.status<400))throw new h("bad-precaching-response",{url:t,status:i.status});i.redirected&&(i=await g(i)),await p({event:s,plugins:r,response:i,request:e===t?a:new Request(e),cacheName:this._cacheName,matchOptions:{ignoreSearch:!0}})}getURLsToCacheKeys(){return this._urlsToCacheKeys}getCachedURLs(){return[...this._urlsToCacheKeys.keys()]}getCacheKeyForURL(e){const t=new URL(e,location.href);return this._urlsToCacheKeys.get(t.href)}async matchPrecache(e){const t=e instanceof Request?e.url:e,n=this.getCacheKeyForURL(t);if(n){return(await self.caches.open(this._cacheName)).match(n)}}createHandler(e=!0){return async({request:t})=>{try{const e=await this.matchPrecache(t);if(e)return e;throw new h("missing-precache-entry",{cacheName:this._cacheName,url:t instanceof Request?t.url:t})}catch(n){if(e)return fetch(t);throw n}}}createHandlerBoundToURL(e,t=!0){if(!this.getCacheKeyForURL(e))throw new h("non-precached-url",{url:e});const n=this.createHandler(t),s=new Request(e);return()=>n({request:s})}}let v;const _=()=>(v||(v=new R),v);const U=(e,t)=>{const n=_().getURLsToCacheKeys();for(const s of function*(e,{ignoreURLParametersMatching:t,directoryIndex:n,cleanURLs:s,urlManipulation:r}={}){const c=new URL(e,location.href);c.hash="",yield c.href;const a=function(e,t=[]){for(const n of[...e.searchParams.keys()])t.some(e=>e.test(n))&&e.searchParams.delete(n);return e}(c,t);if(yield a.href,n&&a.pathname.endsWith("/")){const e=new URL(a.href);e.pathname+=n,yield e.href}if(s){const e=new URL(a.href);e.pathname+=".html",yield e.href}if(r){const e=r({url:c});for(const t of e)yield t.href}}(e,t)){const e=n.get(s);if(e)return e}};let L=!1;function q(e){L||((({ignoreURLParametersMatching:e=[/^utm_/],directoryIndex:t="index.html",cleanURLs:n=!0,urlManipulation:s}={})=>{const r=o();self.addEventListener("fetch",c=>{const a=U(c.request.url,{cleanURLs:n,directoryIndex:t,ignoreURLParametersMatching:e,urlManipulation:s});if(!a)return void 0;let o=self.caches.open(r).then(e=>e.match(a)).then(e=>e||fetch(a));c.respondWith(o)})})(e),L=!0)}const T=e=>{const t=_(),n=r.get();e.waitUntil(t.install({event:e,plugins:n}).catch(e=>{throw e}))},K=e=>{const t=_();e.waitUntil(t.activate())};var b;(function(e){_().addToCacheList(e),e.length>0&&(self.addEventListener("install",T),self.addEventListener("activate",K))})([{'revision':'21fcf393be127eb21f3f3032c1159ad4','url':'./index.html'},{'revision':'2d62fdac3e53a283c6e0228fa7f7f1ab','url':'css/app.css'},{'revision':'f30e029218958bd6aad3cd424f865f94','url':'fonts/Framework7Icons-Regular.eot'},{'revision':'1b6b2c3ed476f4d4b7555af75e387d73','url':'fonts/Framework7Icons-Regular.ttf'},{'revision':'8f897db6f41a6ae1661072172143a21b','url':'fonts/Framework7Icons-Regular.woff'},{'revision':'9393ad14858229d680936a6206688704','url':'fonts/Framework7Icons-Regular.woff2'},{'revision':'e79bfd88537def476913f3ed52f4f4b3','url':'fonts/MaterialIcons-Regular.eot'},{'revision':'a37b0c01c0baf1888ca812cc0508f6e2','url':'fonts/MaterialIcons-Regular.ttf'},{'revision':'012cf6a10129e2275d79d6adac7f3b02','url':'fonts/MaterialIcons-Regular.woff'},{'revision':'570eb83859dc23dd0eec423a49e147fe','url':'fonts/MaterialIcons-Regular.woff2'},{'revision':'3ed9575dcc488c3e3a5bd66620bdf5a4','url':'fonts/OpenSans-Regular.ttf'},{'revision':'ba5cde21eeea0d57ab7efefc99596cce','url':'fonts/OpenSans-SemiBold.ttf'},{'revision':'761abfed0753b923056faf6442de4088','url':'fonts/Ubuntu Mono Nerd Font Complete.ttf'},{'revision':'46f0a3ad815962005753c02c4181aadc','url':'fonts/UbuntuMono-Bold.ttf'},{'revision':'e097b71641f2524e09820d9122b7e910','url':'fonts/UbuntuMono-Regular.ttf'},{'revision':'b325efda5a295da5d493c02d98d970c7','url':'images/Brave32x32.png'},{'revision':'0d0dd47de900c170f759b6f19428d7e6','url':'images/Chrome32x32.png'},{'revision':'92483296ca08a45652e7bad144fc1bbe','url':'images/EnergiLogo-Dark.png'},{'revision':'c6504d0ca81bb0eb9f985454e1195707','url':'images/EnergiLogo-Light.png'},{'revision':'277aaf308826e0d281a6f208fe4f255e','url':'images/Firefox32x32.png'},{'revision':'be18d285bf983564a3542fdff79610d6','url':'images/Git-logo.svg'},{'revision':'ea7a40af921e42df5118ce65dcb14ca1','url':'images/Go-Logo_Aqua.svg'},{'revision':'4baee8feb05d30a5fcd628c6f6a54f9d','url':'images/MicrosoftEdge32x32.png'},{'revision':'4f16c5d909f2f80618d13fc103e6a0dd','url':'images/NRG16x16.png'},{'revision':'05662b8965c45b89968c32fd9b3ad93b','url':'images/Opera32x32.png'},{'revision':'fc5759c17685150e3d4e34619a658630','url':'images/Safari32x32.png'},{'revision':'efe9d37563a8c1710020d29a5316eb07','url':'images/connectedRemote.png'},{'revision':'0f5f879e9febc5f7fdafcb7390ffdc9b','url':'images/connectedRemote.svg'},{'revision':'e53cf675d804be410be9886354215c3c','url':'images/energi_logo.svg'},{'revision':'626bd431e22bfb4897f6f29a23e62ab4','url':'images/ethers_logo.svg'},{'revision':'2d0af7d6135573a78cfbbf94cdee5421','url':'images/favicon.ico'},{'revision':'2d0af7d6135573a78cfbbf94cdee5421','url':'images/favicon.png'},{'revision':'0c0d55564fc4495cdee2a501b7a15cd5','url':'images/framework7-24x24.png'},{'revision':'dba50c88d43fb5e2842a5bda11431878','url':'images/newtux.svg'},{'revision':'8de35400c1a46d910d1238af0fab8120','url':'images/svelte-logo-horizontal.svg'},{'revision':'02c9cc26ade09380cf639cf52a45f7e8','url':'images/svelte_logo.svg'},{'revision':'2c3e5052be6af4fae8638d28f94869ab','url':'images/unknown32x32.png'},{'revision':'eccd81650d8c0a188d65a467ce9fe95c','url':'js/app.js'},{'revision':'b98992ac9c4e0c5dc86c372616b213cb','url':'js/app.js.LICENSE.txt'},{'revision':'440a482350a3b17e414202062c783a75','url':'manifest.json'},{'revision':'afe59015f345e427f8330572620928d1','url':'static/emoji-list.css'},{'revision':'bbcffc64e2471196da5480f40f5872bd','url':'static/icons/128x128.png'},{'revision':'f01c2271d81aa88f829650d18a1b13f9','url':'static/icons/144x144.png'},{'revision':'398bca504224da0d5ee9b04cf8c54ccd','url':'static/icons/152x152.png'},{'revision':'bc19cbf2fbcd28c8c229069e38552d73','url':'static/icons/192x192.png'},{'revision':'cf40c54bf94f558f3e9b2ac29c93d87d','url':'static/icons/256x256.png'},{'revision':'64dad1b90ff9aa6dc27a596fbd8cc9ad','url':'static/icons/512x512.png'},{'revision':'f01c2271d81aa88f829650d18a1b13f9','url':'static/icons/android-chrome-144x144.png'},{'revision':'bc19cbf2fbcd28c8c229069e38552d73','url':'static/icons/android-chrome-192x192.png'},{'revision':'c842a13d002397eecac8e997c779982e','url':'static/icons/android-chrome-36x36.png'},{'revision':'756828018180632663f317de0e949337','url':'static/icons/android-chrome-48x48.png'},{'revision':'a26843458a1874fc68481a0a86ec432e','url':'static/icons/android-chrome-72x72.png'},{'revision':'2d0af7d6135573a78cfbbf94cdee5421','url':'static/icons/android-chrome-96x96.png'},{'revision':'63d53aad0f0eac8a39609a7556ccea22','url':'static/icons/apple-touch-icon-114x114.png'},{'revision':'7eda92c02d504191c58c842cdd5d15db','url':'static/icons/apple-touch-icon-120x120.png'},{'revision':'7482eaac8de980b823f4db58d014c5a3','url':'static/icons/apple-touch-icon-144x144.png'},{'revision':'74636debaad61fd4ab6be95550b3d25d','url':'static/icons/apple-touch-icon-152x152.png'},{'revision':'ffbde802f68369ab4d34d4ea1d50ed46','url':'static/icons/apple-touch-icon-180x180.png'},{'revision':'960454baabc81e55e68f4c5d98389099','url':'static/icons/apple-touch-icon-57x57.png'},{'revision':'72ba79c7dd504d86c21e78c504f2fc1b','url':'static/icons/apple-touch-icon-60x60.png'},{'revision':'3bc7d67645d930d7a0671e84b9ed5d01','url':'static/icons/apple-touch-icon-72x72.png'},{'revision':'6f82c1bfd01fd0f14555134a9cd0dfae','url':'static/icons/apple-touch-icon-76x76.png'},{'revision':'3a0da6068ddf1330bf7feb8194699325','url':'static/icons/apple-touch-icon-precomposed.png'},{'revision':'3a0da6068ddf1330bf7feb8194699325','url':'static/icons/apple-touch-icon.png'},{'revision':'b4e98caaa9798e5bf5395ae16e88a3bc','url':'static/icons/apple-touch-startup-image-1182x2208.png'},{'revision':'e4440e581a51149657be4da3750dc738','url':'static/icons/apple-touch-startup-image-1242x2148.png'},{'revision':'7333f781494231cab6237f6d99918178','url':'static/icons/apple-touch-startup-image-1496x2048.png'},{'revision':'13612082f33e56961524c96c44972b76','url':'static/icons/apple-touch-startup-image-1536x2008.png'},{'revision':'d8ba8cc95ece62892513e6d17c67c814','url':'static/icons/apple-touch-startup-image-320x460.png'},{'revision':'ac5cfac645eec04516367508d51336b2','url':'static/icons/apple-touch-startup-image-640x1096.png'},{'revision':'2cf249abea757fdc788ac2aa37587473','url':'static/icons/apple-touch-startup-image-640x920.png'},{'revision':'da00e1109c54f9afde485643fc27524b','url':'static/icons/apple-touch-startup-image-748x1024.png'},{'revision':'4780f013f094fb24d639a7c6036405dd','url':'static/icons/apple-touch-startup-image-750x1024.png'},{'revision':'3290a70d74fb38686a909880137fe5ce','url':'static/icons/apple-touch-startup-image-750x1294.png'},{'revision':'0f5345958cfb0c95ca42e58bf511ca1b','url':'static/icons/apple-touch-startup-image-768x1004.png'},{'revision':'907efaca20d26e765ba7b267b66f3b6f','url':'static/icons/browserconfig.xml'},{'revision':'d25fbc865be48fc62f766baa3c9e681a','url':'static/icons/example.html'},{'revision':'4f16c5d909f2f80618d13fc103e6a0dd','url':'static/icons/favicon-16x16.png'},{'revision':'79c735d3fa484d563a0c46167638b32e','url':'static/icons/favicon-32x32.png'},{'revision':'2d0af7d6135573a78cfbbf94cdee5421','url':'static/icons/favicon-96x96.png'},{'revision':'2d0af7d6135573a78cfbbf94cdee5421','url':'static/icons/favicon.ico'},{'revision':'2d0af7d6135573a78cfbbf94cdee5421','url':'static/icons/favicon.png'},{'revision':'06f07ba72d8976859c9f7cefb13ead30','url':'static/icons/fonttest.html'},{'revision':'276f1c22d2a52117def754c513db56ff','url':'static/icons/manifest.json'},{'revision':'4800e1902a8523a50a18432b2176190d','url':'static/icons/mstile-150x150.png'},{'revision':'c394ec042cf5f12d4c5d48f0b8ca0640','url':'static/icons/mstile-310x150.png'},{'revision':'f83f3b527c4fc54711c7bfeed0ef4ed2','url':'static/icons/mstile-310x310.png'},{'revision':'25071e18a17a7f3bcd8d512e697c3906','url':'static/icons/mstile-70x70.png'}]||[]),q(b)}]);
//# sourceMappingURL=service-worker.js.map