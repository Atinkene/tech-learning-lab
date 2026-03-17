const User = require('../models/userModel');

const getAllUsers = async (req, res) => {
    const users = await User.getAllUsers();
    return users;
} 

module.exports = { getAllUsers }