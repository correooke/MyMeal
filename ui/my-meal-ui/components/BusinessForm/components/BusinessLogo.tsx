import { CardMedia } from "@mui/material";
import { useTheme } from "@mui/material/styles";

const BusinessLogo = () => {
  const theme = useTheme();
  return (
    <CardMedia
      sx={{
        width: 120,
        height: 120,
        position: "relative",
        display: "inline",
        marginTop: "-60px",
        marginBottom: "-60px",
        borderRadius: theme.border.radius,
        border: `1px solid ${theme.palette.primary.main}`,
      }}
      component="img"
      image={"/resources/images/business-logo.png"}
      alt="Business Logo Image"
    />
  );
};

export default BusinessLogo;
