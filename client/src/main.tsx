import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { TanstackReactQueryProvider } from "./providers/tanstack-react-query.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <TanstackReactQueryProvider>
      <App />
    </TanstackReactQueryProvider>
  </StrictMode>
);
