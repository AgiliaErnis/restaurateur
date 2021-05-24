export default function validateUsername(values) {
    let errors = {};

    if (!values.newUsername.trim()) {
      errors.newUsername = 'Username required';
    } else if (values.newUsername.length < 2) {
      errors.newUsername = 'Username needs to be 2 characters or more'
  }

    return errors;
  }