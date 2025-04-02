<?php
return [
    'cache' => [
        'frontend' => [
            'default' => [
                'id_prefix' => '4bd_'
            ],
            'page_cache' => [
                'id_prefix' => '4bd_'
            ]
        ]
    ],
    'MAGE_MODE' => 'developer',
    'cache_types' => [
        'compiled_config' => 1,
        'config' => 1,
        'layout' => 1,
        'block_html' => 1,
        'collections' => 1,
        'reflection' => 1,
        'db_ddl' => 1,
        'eav' => 1,
        'customer_notification' => 1,
        'config_integration' => 1,
        'config_integration_api' => 1,
        'full_page' => 1,
        'target_rule' => 1,
        'config_webservice' => 1,
        'translate' => 1,
        'vertex' => 1,
        'amasty_shopby' => 1
    ],
    'backend' => [
        'frontName' => 'adminbdb5eeb265f4'
    ],
    'db' => [
        'connection' => [
            'default' => [
                'host' => 'mysql',
                'username' => 'root',
                'dbname' => 'ddvc',
                'password' => 'mysql'
            ],
            'indexer' => [
                'host' => 'mysql',
                'username' => 'root',
                'dbname' => 'ddvc',
                'password' => 'mysql'
            ]
        ]
    ],
    'queue' => [
        'consumers_wait_for_messages' => 0
    ],
    'crypt' => [
        'key' => 'bdb5eeb265f4bff28695f4bb990fe80b'
    ],
    'resource' => [
        'default_setup' => [
            'connection' => 'default'
        ]
    ],
    'x-frame-options' => 'SAMEORIGIN',
    'session' => [
        'save' => 'files'
    ],
    'lock' => [
        'provider' => 'db',
        'config' => [
            'prefix' => null
        ]
    ],
    'downloadable_domains' => [
        'vde.local'
    ],
    'install' => [
        'date' => 'Tue, 04 Aug 2020 05:54:37 +0000'
    ],
    'system' => [
        'default' => [
            'catalog' => [
                'search' => [
                    'engine' => 'elasticsearch7',
                    'elasticsearch7_server_hostname' => 'elasticsearch7',
                    'elasticsearch7_server_port' => 9200
                ]
            ]
        ]
    ]
];
