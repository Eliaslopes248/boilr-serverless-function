/**
 * This file implements the cloudflare service to interact with the registry through cloudflare
 * 
 * @author  Elias Lopes
 * @date    02/20/2026
 * 
 */

const Registry = require("./registry.js");

class CloudflareService extends Registry 
{

    service;
    constructor () {}
    /** methods to override */
    getRegistryMetaData(){}
    
}

module.exports = CloudflareService;