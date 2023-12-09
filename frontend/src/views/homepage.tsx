import React, {useState} from "react";

export const Homepage: React.FC = () => {
    const [title, setTitle] = useState("");
    const [paste, setPaste] = useState("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        const data = { title, paste };

        try {
            const response = await fetch("http://localhost:8080/paste", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });

            if (response.ok) {
                console.log("Paste successfully submitted!");
                // Добавьте здесь код для обработки успешной отправки
            } else {
                console.error("Error submitting paste");
                // Добавьте здесь код для обработки ошибки отправки
            }
        } catch (error) {
            console.error("Error submitting paste", error);
            // Добавьте здесь код для обработки ошибки отправки
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label htmlFor="title">Title:</label>
                <input
                    type="text"
                    id="title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
            </div>
            <div>
                <label htmlFor="paste">Paste:</label>
                <textarea
                    id="paste"
                    rows={5}
                    value={paste}
                    onChange={(e) => setPaste(e.target.value)}
                ></textarea>
            </div>
            <button type="submit">Submit</button>
        </form>
    );
};