import './App.css';
// import axios from 'axios'
// import {useEffect, useState} from "react";
import {Routes, Route, BrowserRouter} from "react-router-dom";
import {PastePage} from "./views/paste";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="" element={<PastePage/>}/>
                <Route path="paste/:id" element={<PastePage/>}/>
            </Routes>
        </BrowserRouter>
    );

}

export default App;