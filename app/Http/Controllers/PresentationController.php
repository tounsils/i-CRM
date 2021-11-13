<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class PresentationController extends Controller
{
    public function index(){
        return view('welcome');
    }

    public function about(){
        return view('about');
    }

    public function LaravelQuestions(){
        return view('LaravelQuestions');
    }
    
    public function LaravelDatabaseSeeding(){
        return view('LaravelDatabaseSeeding');
    }
    public function LaravelTinker(){
        return view('LaravelTinker');
    }
}

