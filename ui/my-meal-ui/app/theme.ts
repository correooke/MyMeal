'use client'; 

declare module '@mui/material/styles' {
  interface Theme {
    header: {
      height: string;
    };
    border: {
      radius: string;
    },
    mainWidth: string;
  }
  // allow configuration using `createTheme`
  interface ThemeOptions {
    header?: {
      height?: string;
    };
    border?: {
      radius?: string;
    },
    mainWidth?: string;
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
    border: {
      radius: "10px",
    },
    mainWidth: "800px",
    components: {
      MuiTextField: {
        styleOverrides: {
          root: {

          },
          
        },
      },
      MuiFilledInput: {
        styleOverrides: {
          root: {
            backgroundColor: "#FFF",
          },
        },
      },


    }

  });

  export default theme;