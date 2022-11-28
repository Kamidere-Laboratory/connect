import {defineConfig} from 'vitepress'

import {
    discordLink,
    gitHubLink,
    additionalTitle,
    commitRef,
    editLink,
   projects
} from '../../shared/'

const ogUrl = 'https://connect.minekube.com'
const ogImage = `${ogUrl}/og-image.png`
const ogTitle = 'Minekube Connect'
const ogDescription = 'Next Generation Minecraft Networks'

export default defineConfig({
    title: `Connect${additionalTitle}`,
    description: ogDescription,

    head: [
        ['link', {rel: 'icon', type: 'image/png', href: '/favicon.png'}],
        ['meta', {property: 'og:type', content: 'website'}],
        ['meta', {property: 'og:title', content: ogTitle}],
        ['meta', {property: 'og:image', content: ogImage}],
        ['meta', {property: 'og:url', content: ogUrl}],
        ['meta', {property: 'og:description', content: ogDescription}],
        ['meta', {name: 'theme-color', content: '#646cff'}],
        // [
        //     'script',
        //     {
        //         src: 'https://cdn.usefathom.com/script.js',
        //         'data-site': 'CBDFBSLI',
        //         'data-spa': 'auto',
        //         defer: ''
        //     }
        // ]
    ],

    vue: {
        reactivityTransform: true
    },

    themeConfig: {
        logo: '/logo.svg',

        editLink: editLink('connect'),

        socialLinks: [
            {icon: 'discord', link: discordLink},
            {icon: 'github', link: `${gitHubLink}/connect-java`}
        ],

        // algolia: {
        //     appId: '7H67QR5P0A',
        //     apiKey: 'deaab78bcdfe96b599497d25acc6460e',
        //     indexName: 'vitejs',
        //     searchParameters: {
        //         facetFilters: ['tags:en']
        //     }
        // },

        // carbonAds: {
        //     code: 'CEBIEK3N',
        //     placement: 'vitejsdev'
        // },

        footer: {
            message: `Plugins are released under the MIT License. (web version: ${commitRef})`,
            copyright: 'Copyright © 2022-present Minekube & Contributors'
        },

        nav: [
            {text: 'Guide', link: '/guide/'},
            {text: 'Quick Start', link: '/guide/quick-start'},
            {text: 'Downloads', link: '/guide/downloads'},
            projects,
        ],

        sidebar: {
            '/guide/': [
                {
                    text: 'Getting Started',
                    items: [
                        {text: 'Introduction', link: '/guide/'},
                        {text: 'Quick Start', link: '/guide/quick-start'},
                        {text: 'Downloads', link: '/guide/downloads'},
                        {text: 'Why', link: '/guide/why'},
                    ]
                },
                {
                    text: 'Guide',
                    items: [
                        {
                            text: 'Advertising your Server',
                            link: '/guide/advertising-your-server'
                        },
                        {
                            text: 'Custom Domains',
                            link: '/guide/custom-domains'
                        },
                        {
                            text: 'DDoS & Bot protection',
                            link: '/guide/protections'
                        },
                    ]
                },
                {
                    text: 'Roadmap',
                    items: [
                        {
                            text: 'Development & Use cases',
                            link: '/guide/roadmap'
                        },
                    ]
                },
                // {
                //     text: 'APIs',
                //     items: [
                //         {
                //             text: 'Plugin API',
                //             link: '/guide/api-plugin'
                //         },
                //     ]
                // }
            ],
        }
    }
})
