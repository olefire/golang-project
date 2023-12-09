import './App.css';
// import axios from 'axios'
// import {useEffect, useState} from "react";
import {Routes, Route, BrowserRouter} from "react-router-dom";
import {Homepage} from "./views/homepage";
import {Batch} from "./views/batch";
import {PastePage} from "./views/paste";
import {TestLintPage} from "./views/testLintPage";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="main" element={<Homepage/>}/>
                <Route path="batch" element={<Batch/>}/>
                <Route path="paste/:id" element={<PastePage/>}/>
                <Route path="lint" element={<TestLintPage/>}/>
            </Routes>
        </BrowserRouter>
    );

}

export default App;