import React, {useContext} from 'react';
import zxcvbn from 'zxcvbn';
import "./PasswordStrengthMeter.css"
import { UserContext } from "../../../UserContext"

const PasswordStrengthMeter = ({ password }) => {
    const { setGoodPassword } = useContext(UserContext)
    const testResult = zxcvbn(password);
    const num = testResult.score * 100 / 3;

    const createPassLabel = () => {
        if (testResult.score > 1) {
          setGoodPassword(true)
        } else {
            setGoodPassword(false)
        }
        if (num <= 100) {
            switch (testResult.score) {

                case 0:
                    return 'Short';
                case 1:
                    return 'Weak';
                case 2:
                    setGoodPassword(true);
                    return 'Good';
                case 3:
                    return 'Strong';
                default:
                    return '';
            }
        }
        else if (password.length > 64) {
            return "Too Long";
        } else {
                return "Strong";
            }
    }

    const funcProgressColor = () => {
        if (num <= 100) {
            switch (testResult.score) {
                case 0:
                    return '#828282';
                case 1:
                    return '#EA1111';
                case 2:
                    return '#9bc158';
                case 3:
                    return '#00b500';
                default:
                    return 'none';
            }
        } else if (password.length > 64) {
            return "#EA1111"
        } else {
            return '#00b500';
        }
    }

  const changePasswordColor = () => ({
    width: `${num <= 100 ? num : 100}%`,
    background: funcProgressColor(),
    height: '5px'
  })

  return (
    <>
      <div className="progress" style={{ height: '5px' }}>
        <div className="progress-bar" style={changePasswordColor()}></div>
      </div>
      <p style={{ color: funcProgressColor(), fontSize:"12px"}}>{createPassLabel()}</p>
    </>
  )
}

export default PasswordStrengthMeter