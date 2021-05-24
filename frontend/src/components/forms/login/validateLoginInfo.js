export default function validateLoginInfo(values) {
    let errors = {};

    if (!values.email) {
      errors.email = 'Email required';
    }

    if (!values.password) {
      errors.password = 'Password is required';
    }

    return errors;
  }