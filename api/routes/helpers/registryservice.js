/**
 * This file will implement the RegistryService class that interacts with 
 * the registry holding builds
 * 
 * @author  Elias Lopes
 * @date    02/20/2026
 * 
 */

/** includes */
const CloudflareService = require("./registry-services/cloudflareservice");


/**
 * defines the RegistryService class
 * @returns RegistryService
 */
class RegistryService 
{
    /** services */
    service;

    /** methods */
    constructor(service=null) { this.service = service; }

    /** set the registry type with dependency injection */
    setRegistryService(
        service)
    {   
        if (!service instanceof Registry){ return; }
        this.service = service;
    }
    
}

