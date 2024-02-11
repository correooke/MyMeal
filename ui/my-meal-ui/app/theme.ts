'use client'; 

declare module '@mui/material/styles' {
  interface Theme {
    header: {
      height: string;
    };
  }
  // allow configuration using `createTheme`
  interface ThemeOptions {
    header?: {
      height?: string;
    };
  }
}

import { createTheme } from "@mui/material/styles";

const theme = createTheme({
    palette: {
      primary: {
        main: "#6C63FF",
        contrastText: "#FFF",
      },
      secondary: {
        main: "#FFF",
        contrastText: "#6C63FF",
      },
    },
    header: {
      height: "130px",
    },
  });

  export default theme;