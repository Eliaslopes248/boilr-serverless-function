/**
 * 
 * This file is the entry point of the serverless api
 * that will handle http business logic for boilr
 * 
 * @author Elias Lopes
 * @date   02/20/2026
 * 
 */


const express   = require("express");
const app       = express();
app.use(express.json());

/**
 * @import all route modules from different files to be added to app instance
 */

const healthCheckRoutes = require("routes/tests/healthcheck.js");


/**
 * @use route modules
 */
app.use("/api/tests", healthCheckRoutes);


// export for vercels use
module.exports = app;

