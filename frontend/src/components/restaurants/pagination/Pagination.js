import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Pagination from '@material-ui/lab/Pagination';
import './Pagination.css'

const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
        margin: '4rem',
    },
  },
}));

export default function RestaurantPagination({ restaurantsPerPage, totalRestaurants, paginate }) {

  const pageNumbers = [];
  for (let i = 1; i <= Math.ceil(totalRestaurants / restaurantsPerPage); i++) {
    pageNumbers.push(i);
  }

  const classes = useStyles();

    return (
      <div className={classes.root}>
        <Pagination
          count={pageNumbers.length}
          shape="rounded"
          size="large"
          showFirstButton
          showLastButton
          onChange={(e, value) => paginate(value)}
        />
        </div>
    )
}