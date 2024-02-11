import { CardMedia } from "@mui/material";

const BusinessLogo = () => {
  return (
    <CardMedia
      sx={{
        width: 120,
        height: 120,
        position: "relative",
        marginTop: "-60px",
        padding: 0,
        borderRadius: "5%",
        border: "1px solid blue",
      }}
      component="img"
      image={"/resources/images/business-logo.png"}
      alt="Business Logo Image"
    />
  );
};

export default BusinessLogo;
