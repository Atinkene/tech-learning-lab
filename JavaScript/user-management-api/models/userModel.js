const db = require('../config/db')
const bcrypt = require('bcryptjs')


class User {
    constructor ({id, username, email, firstname, lastname, password, role}) {
        this.id =  id;
        this.username = username;
        this.email = email;
        this.firsname = firstname;
        this.lastname = lastname;
        this.password = password;
        this.role = role
    }

    static async getAllUsers() {
        const [result] = await db.execute(
            'SELECT id, username, email, firstname, lastname, role FROM users'
        );
        return result;
    }result

    static async getUserById ( id ) {
        const [rows] = await db.execute(
            'SELECT * FROM users WHERE id = ?', id
        );
        return rows[0];
    }
}

module.exports = User;
