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
| in laravel Route there is two parameters url and call back 
|
*/
/*
Route::get('/', function () {
    return view('welcome');
});
*/

Route::get('/', 'App\Http\Controllers\PresentationController@index');
Route::get('/about', 'App\Http\Controllers\PresentationController@about');
Route::get('/LaravelQuestions', 'App\Http\Controllers\PresentationController@LaravelQuestions');
Route::get('/LaravelDatabaseSeeding', 'App\Http\Controllers\PresentationController@LaravelDatabaseSeeding');

Route::get('/blog', 'App\Http\Controllers\BlogController@index');

Route::get('/hello', function () {
    return "This is a simple Route returning hello message";
});

Route::get('users/{id}', function ($id) {
    return "Your id id $id";
});

Route::get('/users', 'App\Http\Controllers\usersController@index')->name('home');
Route::get('/createuser', 'App\Http\Controllers\usersController@create')->name('create');
Route::POST('/createuser', 'App\Http\Controllers\usersController@store')->name('storeuser');

