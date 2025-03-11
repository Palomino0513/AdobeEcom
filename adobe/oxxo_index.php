<?php
/**
 * Public alias for the application entry point
 *
 * Copyright © Magento, Inc. All rights reserved.
 * See COPYING.txt for license details.
 */

use Magento\Framework\App\Bootstrap;

try {
    require __DIR__ . '/../app/bootstrap.php';

    function isHttpHost(string $host) {
        if (!isset($_SERVER['HTTP_HOST'])) {
            return false;
        }
        return $_SERVER['HTTP_HOST'] === $host;
    }
    if (isHttpHost("oxxoclub.local")) {
        $_SERVER["MAGE_RUN_CODE"] = "oxxo_pronto";
        $_SERVER["MAGE_RUN_TYPE"] = "website";
    }
    if (isHttpHost("oxxobusiness.local")) {
        $_SERVER["MAGE_RUN_CODE"] = "oxxo_club_b2b2b";
        $_SERVER["MAGE_RUN_TYPE"] = "website";
    }
} catch (\Exception $e) {
    echo <<<HTML
<div style="font:12px/1.35em arial, helvetica, sans-serif;">
    <div style="margin:0 0 25px 0; border-bottom:1px solid #ccc;">
        <h3 style="margin:0;font-size:1.7em;font-weight:normal;text-transform:none;text-align:left;color:#2f2f2f;">
        Autoload error</h3>
    </div>
    <p>{$e->getMessage()}</p>
</div>
HTML;
    http_response_code(500);
    exit(1);
}

$bootstrap = Bootstrap::create(BP, $_SERVER);
/** @var \Magento\Framework\App\Http $app */
$app = $bootstrap->createApplication(\Magento\Framework\App\Http::class);
$bootstrap->run($app);
