/**
 * 
 * This file is made to test the api to ensure reliable endpoints
 * 
 * @author Elias Lopes
 * @date   02/20/2026
 * 
 */

const express   = require("express");
const router    = express.Router();

/** simple GET request to make sure communication is established */
router.get(
    "/health", 
    (req, res)=>{
        res.status(200).json({
            status:200,
            message: "API GET is healthy"
        });
});

router.post(
    "/health", 
    (req, res)=>{
        res.status(200).json({
            status:     200,
            message:    "API POST is healthy",
            body:       req?.body
        });
});

module.exports = router;