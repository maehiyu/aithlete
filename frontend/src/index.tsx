import React from "react";
import ReactDOM from "react-dom/client";
import App from "./pages/App";
import { Amplify } from 'aws-amplify';
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import "./index.css";
import reportWebVitals from "./reportWebVitals";
import { BrowserRouter } from "react-router-dom";
import { Authenticator } from "@aws-amplify/ui-react";

Amplify.configure({
  Auth: {
    Cognito: {
      userPoolId: process.env.REACT_APP_AWS_COGNITO_USER_POOLS_ID || '',
      userPoolClientId: process.env.REACT_APP_AWS_COGNITO_USER_POOLS_CLIENT_ID || '',
      // identityPoolId: process.env.REACT_APP_AWS_COGNITO_IDENTITY_POOL_ID || '',
    }
  }
});

const queryClient = new QueryClient();

const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement);
root.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Authenticator loginMechanisms={["email"]}>
          <App />
        </Authenticator>
      </BrowserRouter>
    </QueryClientProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
