import React, { useEffect, useState } from "react";
import API from "../api/api";

function Admin() {
    const [requests, setRequests] = useState([]);

    useEffect(() => {
        const loadRequests = async () => {
            try {
                const res = await API.get("/admin/whitelist-requests");
                setRequests(res.data);
            } catch (err) {
                console.error(err);
            }
        };

        loadRequests();
    }, []);

    const approve = async (id) => {
        await API.post("/admin/whitelist/approve", { user_id: id });
        window.location.reload();
    };

    const reject = async (id) => {
        await API.post("/admin/whitelist/reject", { user_id: id });
    };

    return (
        <div>
            <h2>Admin Panel</h2>

            {requests.map((r) => (
                <div key={r.id}>
                    {r.minecraft_username}

                    <button onClick={() => approve(r.id)}>Approve</button>
                    <button onClick={() => reject(r.id)}>Reject</button>
                </div>
            ))}
        </div>
    );
}

export default Admin;