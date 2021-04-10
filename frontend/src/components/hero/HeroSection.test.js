import * as React from "react";
import * as ReactDOM from "react-dom"
import HeroSection from './HeroSection'

test('testing', () => {
    const root = document.createElement("div");
    ReactDOM.render(<HeroSection />, root);
    expect(root.querySelector("h1").textContent).toBe("Discover the best dining destinations in Prague");
})
