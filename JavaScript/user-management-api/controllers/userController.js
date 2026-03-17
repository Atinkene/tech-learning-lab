const User = require("../models/userModel")
const userService = require('../services/userService')

const getAllUsers = async (req, res) => {
    const users = await userService.getAllUsers();
    return res.json(users);
}

module.exports = { getAllUsers }