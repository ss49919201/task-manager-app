import React, { useEffect, useState } from 'react';

type User = {
    id: number;
    name: string;
}

const getUserList = async (): Promise<User[]> => {
    const res = await fetch('http://127.0.0.1:12345/users', {
        method: 'GET',
    });
    return await res.json();
}

const App = () => {
    const [users, setUsers] = useState<User[]>([]);
    useEffect(() => {
        getUserList()
            .then(
                (result) => {
                    setUsers(result);
                }
            )
    })

    return (
        <>
            <ul>
                {
                    users.map((user) => <li>{user.name}</li>)
                }
            </ul>
        </>
    );
}

export default App;
