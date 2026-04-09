/**
 * This file defines a interface for registry services,
 * will be used as base class for real services
 * 
 * @author  Elias Lopes
 * @date    02/20/2026
 * 
 */

const { error } = require("console");


class Registry 
{
    constructor () {}
    /** abstract methods to override */
    getRegistryMetaData(){ throw error("Cannot call method through interface");}

}

module.exports = Registry;