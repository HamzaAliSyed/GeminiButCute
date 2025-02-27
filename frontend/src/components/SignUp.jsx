import { FaUserPlus } from "react-icons/fa";
import { useState } from "react";

const SignUp = () => {
  const [userName, setUserName] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const [message, setMessage] = useState("");
  const [passwordError, setPasswordError] = useState("");
  const [showPasswordRules, setShowPasswordRules] = useState(false);

  const passwordRules = [
    "Minimum 13 characters",
    "At least one uppercase letter",
    "At least one lowercase letter",
    "At least one number",
    "At least one special character"
  ];

  const checkPasswordValidity = (pwd) => {
    const isLongEnough = pwd.length >= 13;
    const hasUpper = /[A-Z]/.test(pwd);
    const hasLower = /[a-z]/.test(pwd);
    const hasNumber = /[0-9]/.test(pwd);
    const hasSpecial = /[!@#$%^&*(),.?":{}|<>]/.test(pwd);
    return isLongEnough && hasUpper && hasLower && hasNumber && hasSpecial;
  };

  const hashPassword = async (pwd) => {
    const encoder = new TextEncoder();
    const data = encoder.encode(pwd);
    const hashBuffer = await window.crypto.subtle.digest("SHA-256", data);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, "0")).join("");
    return hashHex;
  };

  const userRegistrationHandle = async (e) => {
    e.preventDefault();

    if (password !== confirmPassword) {
      setPasswordError("Passwords do not match");
      setTimeout(() => setPasswordError(""), 3000);
      return;
    }

    if (!checkPasswordValidity(password)) {
      setPasswordError("Password does not meet the required rules");
      setTimeout(() => setPasswordError(""), 3000);
      return;
    }

    const hashedPassword = await hashPassword(password);

    const payload = {
      username: userName,
      firstName,
      lastName,
      email,
      password: hashedPassword,
    };

    try {
      const response = await fetch("http://localhost:2712/v1/user/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (response.ok) {
        setMessage("User registered successfully!");
        alert("User registered successfully!");
      } else {
        const errorData = await response.text();
        setMessage(`Registration failed: ${errorData}`);
        alert(`Registration failed: ${errorData}`);
      }
    } catch (error) {
      setMessage(`An error occurred: ${error.message}`);
      alert(`An error occurred: ${error.message}`);
    }
  };

  return (
    <div className="flex flex-col justify-center items-center h-[100vh] background-image">
      <div className="bg-white shadow-lg p-5 rounded-xl h-[33rem] w-[20rem] flex flex-col justify-center items-center">
        <div className="mb-10 mt-5">
          <h1 className="text-center font-bold text-[20px]">Sign Up</h1>
          <p className="text-center text-sm text-gray-400">
            Register for{" "}
            <span className="font-bold text-gray-500">Gemini But</span>{" "}
            <span className="font-bold italic text-gray-600">Cute</span>
          </p>
        </div>
        <form className="w-full" onSubmit={userRegistrationHandle}>
          <div className="w-full">
            <input
              className="w-full border border-gray-300 p-2 rounded-md bg-[#f1f1f1be] text-[#383838f3] mb-3 font-medium outline-none placeholder:text-[#606060f3]"
              type="text"
              id="username"
              name="username"
              autoComplete="username"
              placeholder="Username"
              value={userName}
              onChange={(e) => setUserName(e.target.value)}
              required
            />
            <input
              className="w-full border border-gray-300 p-2 rounded-md bg-[#f1f1f1be] text-[#383838f3] mb-3 font-medium outline-none placeholder:text-[#606060f3]"
              type="text"
              id="firstname"
              name="firstname"
              placeholder="First Name"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
              required
            />
            <input
              className="w-full border border-gray-300 p-2 rounded-md bg-[#f1f1f1be] text-[#383838f3] mb-3 font-medium outline-none placeholder:text-[#606060f3]"
              type="text"
              id="lastname"
              name="lastname"
              placeholder="Last Name"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
              required
            />
            <input
              className="w-full border border-gray-300 p-2 rounded-md bg-[#f1f1f1be] text-[#383838f3] mb-3 font-medium outline-none placeholder:text-[#606060f3]"
              type="email"
              id="email"
              name="email"
              autoComplete="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
            <div className="relative">
              <input
                className="w-full border border-gray-300 p-2 rounded-md bg-[#f1f1f1be] text-[#383838f3] mb-3 font-medium outline-none placeholder:text-[#606060f3]"
                type="password"
                id="password"
                name="password"
                placeholder="Password"
                value={password}
                onFocus={() => {
                  setShowPasswordRules(true);
                  setTimeout(() => setShowPasswordRules(false), 3000);
                }}
                onChange={(e) => setPassword(e.target.value)}
                required
              />
              {showPasswordRules && (
                <div className="absolute top-full left-0 bg-gray-200 p-2 rounded-md text-xs shadow-md mt-1">
                  {passwordRules.map((rule, index) => (
                    <div key={index}>{rule}</div>
                  ))}
                </div>
              )}
            </div>
            <input
              className="w-full border border-gray-300 p-2 rounded-md bg-[#f1f1f1be] text-[#383838f3] mb-3 font-medium outline-none placeholder:text-[#606060f3]"
              type="password"
              id="confirmpassword"
              name="confirmpassword"
              placeholder="Confirm Password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              required
            />
          </div>
          {passwordError && (
            <div className="mb-3 text-red-500 text-xs">{passwordError}</div>
          )}
          <div className="w-full my-3">
            <button
              className="text-white bg-gray-800 mb-8 w-full p-2 rounded-md font-bold flex justify-center items-center gap-2 cursor-pointer"
              type="submit"
              disabled={password !== confirmPassword || password === ""}
            >
              Register <FaUserPlus />
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default SignUp;
