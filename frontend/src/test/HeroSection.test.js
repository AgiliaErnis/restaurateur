import * as React from "react";
import * as ReactDOM from "react-dom"
import HeroSection from '../components/hero/HeroSection'

test('hero prints right string', () => {
    const root = document.createElement("div");
    ReactDOM.render(<HeroSection />, root);
    expect(root.querySelector("h1").textContent).toBe("Discover the best dining destinations in Prague");
})
