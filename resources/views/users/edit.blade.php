@extends('layouts.main')

@section('title', 'Users Edit')

@section('content')

<main id="main">

    <!-- ======= company Section ======= -->
    <section id="user" class="user">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Users</h2>
          <p>Listing</p>
        </header>

        <div class="container">
        
        @if (session('message'))
        <div class="alert session('alert-type')">
         {!! session('message') !!}
        </div>
        @endif


<div class="row">
    <div class="col-sm-8 offset-sm-2">
        <h1 class="display-3">Update a user</h1>

        @if ($errors->any())
        <div class="alert alert-danger">
            <ul>
                @foreach ($errors->all() as $error)
                <li>{{ $error }}</li>
                @endforeach
            </ul>
        </div>
        <br /> 
        @endif
        <form method="post" action="{{ route('users.update', $user->id) }}">
            @method('PATCH') 
            @csrf
            <div class="form-group">

                <label for="first_name">First Name:</label>
                <input type="text" class="form-control" name="first_name" value="{{ $user->first_name }}" />
            </div>

            <div class="form-group">
                <label for="last_name">Last Name:</label>
                <input type="text" class="form-control" name="last_name" value="{{ $user->last_name }}" />
            </div>

            <div class="form-group">
                <label for="email">Email:</label>
                <input type="text" class="form-control" name="email" value="{{ $user->email }}" />
            </div>
            <div class="form-group">
                <label for="phone">Phone:</label>
                <input type="text" class="form-control" name="phone" value="{{ $user->phone }}" />
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" class="form-control" name="password" value="" />
            </div>

            <button type="submit" class="btn btn-primary">Update</button>
        </form>
    </div>
</div>

<div class="col-sm-12">

  @if(session()->get('success'))
    <div class="alert alert-success">
      {{ session()->get('success') }}  
    </div>
  @endif
</div>

@endsection