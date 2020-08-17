
import HomePage from '../pages/home.svelte';
import InfoPage from '../pages/info.svelte';
import Web3ProviderPage from '../pages/web3provider.svelte';
import SmartContractPage from '../pages/smartcontract.svelte';


import AboutPage from '../pages/about.svelte';
import FormPage from '../pages/form.svelte';

import LeftPage1 from '../pages/left-page-1.svelte';
import LeftPage2 from '../pages/left-page-2.svelte';
import DynamicRoutePage from '../pages/dynamic-route.svelte';
import RequestAndLoad from '../pages/request-and-load.svelte';
import NotFoundPage from '../pages/404.svelte';

let routes = [
  {
    path: '/',
    component: HomePage,
  },
  {
    path: '/info/',
    component: InfoPage,
  },
  {
    path: '/web3provider/',
    component: Web3ProviderPage,
  },
  {
    path: '/smartcontract/',
    component: SmartContractPage,
  },
  {
    path: '/about/',
    component: AboutPage,
  },
  {
    path: '/form/',
    component: FormPage,
  },

  {
    path: '/left-page-1/',
    component: LeftPage1,
  },
  {
    path: '/left-page-2/',
    component: LeftPage2,
  },
  {
    path: '/dynamic-route/blog/:blogId/post/:postId/',
    component: DynamicRoutePage,
  },
  {
    path: '/request-and-load/user/:userId/',
    async: function (routeTo, routeFrom, resolve, reject) {
      // Router instance
      let router = this;

      // App instance
      let app = router.app;

      // Show Preloader
      app.preloader.show();

      // User ID from request
      let userId = routeTo.params.userId;

      // Simulate Ajax Request
      setTimeout(function () {
        // We got user data from request
        let user = {
          firstName: 'Vladimir',
          lastName: 'Kharlampidi',
          about: 'Hello, i am creator of Framework7! Hope you like it!',
          links: [
            {
              title: 'Framework7 Website',
              url: 'http://framework7.io',
            },
            {
              title: 'Framework7 Forum',
              url: 'http://forum.framework7.io',
            },
          ]
        };
        // Hide Preloader
        app.preloader.hide();

        // Resolve route to load page
        resolve(
          {
            component: RequestAndLoad,
          },
          {
            context: {
              user: user,
            }
          }
        );
      }, 1000);
    },
  },
  {
    path: '(.*)',
    component: NotFoundPage,
  },
];

export default routes;
