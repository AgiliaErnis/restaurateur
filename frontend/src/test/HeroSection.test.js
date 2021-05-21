import * as React from "react";
import * as ReactDOM from "react-dom"
import { BrowserRouter as Router } from 'react-router-dom'
import HeroSection from '../components/hero/HeroSection'
import { UserContext } from "../UserContext";

test('hero prints right string', () => {
    const root = document.createElement("div");
    ReactDOM.render(<UserContext.Provider value={false}><Router><HeroSection /></Router></UserContext.Provider>, root);
    expect(root.querySelector("h1").textContent).toBe("Discover the best dining destinations in Prague");
})
