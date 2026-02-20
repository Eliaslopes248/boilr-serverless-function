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
const dotenv    = require("dotenv");
const app       = express();

/** @middlewares */
app.use(express.json());
dotenv.config();

/**
 * @import all route modules from different files to be added to app instance
 */

const healthCheckRoutes = require("./routes/tests/healthcheck.js");


/**
 * @use route modules
 */
app.use("/api/tests", healthCheckRoutes);

app.get("/", (req,res)=>{
    res.json({
        status: 200,
        message: "hello serverless"
    });
})

const ENV_TYPE = process.env.ENV_TYPE || null;

// only run listener() in dev env
// if (!ENV_TYPE || ENV_TYPE == "dev"){
//     app.listen(3000, ()=>{})
// }


// Export for Vercel (docs: export the Express app)
module.exports = app;

