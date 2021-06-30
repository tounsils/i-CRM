<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/
/*
Route::get('/', function () {
    return view('welcome');
});
*/

Route::get('/', 'App\Http\Controllers\PresentationController@index');
Route::get('/about', 'App\Http\Controllers\PresentationController@about');
Route::get('/blog', 'App\Http\Controllers\BlogController@index');

Route::get('/hello', function () {
    return "This is a simple Route returning hello message";
});

Route::get('users/{id}', function ($id) {
    return "Your id id $id";
});

Route::get('/createuser', 'App\Http\Controllers\user@create');
Route::post('/createuser', 'App\Http\Controllers\user@store');