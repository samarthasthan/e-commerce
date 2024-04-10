import styles from "./MainHeader.module.css";
import SimpleButton from "../../buttons/SimpleButton";
import TextInput from "../../inputs/textInput/TextInput";
import React from "react";
import { NavLink } from "react-router-dom";
function MainHeader() {
  return (
    <div className={styles.mainHeader}>
      <div className={styles.options}>
        <SimpleButton
          backgroundColor="var(--white-color)"
          color="var(--black--color)"
        >
          FruBay
        </SimpleButton>

        <ul className={styles.tabs}>
          <NavLink className={({ isActive, isPending }) =>
            isPending ? "pending" : isActive ? "active" : ""
          } to="/"><li>Dashboard</li></NavLink>
          <NavLink className={({ isActive, isPending }) =>
            isPending ? "pending" : isActive ? "active" : ""
          } to="/sellers"><li>Seller</li></NavLink>
          <li>Products</li>
          <li>Learn</li>
        </ul>
      </div>

      <div className={styles.user}>
        <TextInput>Search...</TextInput>
        <img
          src="https://cdnstorage.sendbig.com/unreal/female.webp"
          alt="user"
          srcset=""
          className={styles.profilePicture}
        />
      </div>
    </div>
  );
}

export default MainHeader;