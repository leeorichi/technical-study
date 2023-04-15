First create global variables (in functions.php or as a mu-plugin):

``` php
<?php

/*
 * CUSTOM GLOBAL VARIABLES
 */
function wtnerd_global_vars() {

	global $wtnerd;
	$wtnerd = array(

		'edition'  => get_query_var('category_name'),
		'channel'  => get_query_var('channel'),
		'tag'      => get_query_var('tag'),

	);

}
add_action( 'parse_query', 'wtnerd_global_vars' );
```

Why are we using an associative array variable `$wtnerd`? Because global variables need to be unique, and by keeping `$wtnerd` unique we can have simpler names for all the variables in its array.

By the way, the same can also be done like this:

``` php
<?php

/*
 * CUSTOM GLOBAL VARIABLES
 */
function wtnerd_global_vars() {

	global $wtnerd;
	$wtnerd['edition'] = get_query_var('category_name');
	$wtnerd['channel'] = get_query_var('channel');
	$wtnerd['tag']     = get_query_var('tag');

}
add_action( 'parse_query', 'wtnerd_global_vars' );
```

Then use `$GLOBALS[];` to call the variable elsewhere (another file):

``` php
<?php

if( $GLOBALS['wtnerd']['edition'] == uk ) {

	// Do something

}
```

***

If the function in which you are defining the global variables is not hooked into a filter or action, e.g. `add_action( 'parse_query', 'wtnerd_global_vars' );` as we are doing above, then you should do it as shown below.

In functions.php or mu-plugin:

``` php
<?php

/*
 * CUSTOM GLOBAL VARIABLES
 */
function wtnerd_global_vars() {

	global $wtnerd;
	$wtnerd = array(

		'edition'  => get_query_var('category_name'),
		'channel'  => get_query_var('channel'),
		'tag'      => get_query_var('tag'),

	);

}
```

Then, to call the variable elsewhere (another file), you need to manually initialize the function before you can use the variable:

``` php
<?php

wtnerd_global_vars();
if( $GLOBALS['wtnerd']['edition'] == uk ) {

	// Do something

}
```
