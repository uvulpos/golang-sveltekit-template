export const prerender = false;
export const ssr = false;
export const trailingSlash = 'always';

import { addMessages, getLocaleFromNavigator, init } from 'svelte-i18n';

import en from './../lib/i18n/en.json';
import de from './../lib/i18n/de.json';

addMessages('en', en);
addMessages('de', de);

init({
    fallbackLocale: 'en',
    // initialLocale: 'en',
    initialLocale: getLocaleFromNavigator(),
  });