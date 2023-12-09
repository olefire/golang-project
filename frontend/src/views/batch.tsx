import axios from 'axios'
import React, {useEffect, useState} from "react";

export const Batch: React.FC = () => {
    const apiURL = "http://localhost:8080/batch";

    const [batch, setBatch] = useState<any[]>([])


    useEffect(() => {
        axios
            .get(apiURL)
            .then(data => {
                setBatch(data.data.data)
            })
    }, []);

    return (
        <div>
            <h1>Данные с API</h1>
            <ul>
                {batch.map(item => (
                    <li key={item.id}>{item.title}: {item.paste}</li>
                ))}
            </ul>
        </div>
    );
}