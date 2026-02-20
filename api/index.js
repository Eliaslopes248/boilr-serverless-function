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

const healthCheckRoutes = require("./routes/tests/healthcheck.js");


/**
 * @use route modules
 */
app.use("/api/tests", healthCheckRoutes);

// Root route so /api and /api/ hit the app (rewrite sends all /api/* here)
app.get("/api", (req, res) => res.status(200).json({ ok: true, message: "API root" }));
app.get("/api/", (req, res) => res.status(200).json({ ok: true, message: "API root" }));

// 404 for unmatched routes
app.use((req, res) => res.status(404).json({ error: "Not found", path: req.url }));

// Export for Vercel (docs: export the Express app)
module.exports = app;

