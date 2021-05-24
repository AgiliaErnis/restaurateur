export default function validateDeleteAccount(values) {
    let errors = {};

    if (!values.password) {
        errors.password = 'Password is required'
    }

    return errors;
  }